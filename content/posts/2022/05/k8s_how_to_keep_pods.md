# k8s是怎么维持pod的运行的呢？

当接收了yaml配置的信息后，是怎么维持pod根据声明一直运行的呢？

让我们沿着命令执行的过程来一睹为快：`kubectl apply -f pod.yaml`.

源码位置：`cmd/kubectl/kubectl.go` -> `staging/src/k8s.io/kubectl/pkg/cmd/cmd.go` -> `staging/src/k8s.io/kubectl/pkg/cmd/apply/apply.go`

最终的执行方法：

```go
func (o *ApplyOptions) Run() error {
    // 预处理
	if o.PreProcessorFn != nil {
		klog.V(4).Infof("Running apply pre-processor function")
		if err := o.PreProcessorFn(); err != nil {
			return err
		}
	}

	// Enforce CLI specified namespace on server request.
	if o.EnforceNamespace {
		o.VisitedNamespaces.Insert(o.Namespace)
	}

	// Generates the objects using the resource builder if they have not
	// already been stored by calling "SetObjects()" in the pre-processor.
	errs := []error{}
	infos, err := o.GetObjects()
	if err != nil {
		errs = append(errs, err)
	}
	if len(infos) == 0 && len(errs) == 0 {
		return fmt.Errorf("no objects passed to apply")
	}
	// Iterate through all objects, applying each one.
	for _, info := range infos {
		if err := o.applyOneObject(info); err != nil {
			errs = append(errs, err)
		}
	}
	// If any errors occurred during apply, then return error (or
	// aggregate of errors).
	if len(errs) == 1 {
		return errs[0]
	}
	if len(errs) > 1 {
		return utilerrors.NewAggregate(errs)
	}

	if o.PostProcessorFn != nil {
		klog.V(4).Infof("Running apply post-processor function")
		if err := o.PostProcessorFn(); err != nil {
			return err
		}
	}

	return nil
}

// applyOneObject里会调用以下方法
func (m *Helper) Patch(namespace, name string, pt types.PatchType, data []byte, options *metav1.PatchOptions) (runtime.Object, error) {
	if options == nil {
		options = &metav1.PatchOptions{}
	}
	if m.ServerDryRun {
		options.DryRun = []string{metav1.DryRunAll}
	}
	if m.FieldManager != "" {
		options.FieldManager = m.FieldManager
	}
	if m.FieldValidation != "" {
		options.FieldValidation = m.FieldValidation
	}
	return m.RESTClient.Patch(pt).
		NamespaceIfScoped(namespace, m.NamespaceScoped).
		Resource(m.Resource).
		Name(name).
		SubResource(m.Subresource).
		VersionedParams(options, metav1.ParameterCodec).
		Body(data).
		Do(context.TODO()). // 调用api，把apply请求发到主节点，记录信息到etcd之后，再创建出相应的pod
		Get()
}

// 那么，接收并处理这个Patch请求的代码在哪里呢？

// NewStreamWatcher creates a StreamWatcher from the given decoder.
func NewStreamWatcher(d Decoder, r Reporter) *StreamWatcher {
	sw := &StreamWatcher{
		source:   d,
		reporter: r,
		// It's easy for a consumer to add buffering via an extra
		// goroutine/channel, but impossible for them to remove it,
		// so nonbuffered is better.
		result: make(chan Event),
		// If the watcher is externally stopped there is no receiver anymore
		// and the send operations on the result channel, especially the
		// error reporting might block forever.
		// Therefore a dedicated stop channel is used to resolve this blocking.
		done: make(chan struct{}),
	}
	go sw.receive() // 接收请求，然后通过chan发送出去，再由其它代码来处理？
	return sw
}

// TODO:
```

[apimachinery共享库](https://github.com/kubernetes/apimachinery)

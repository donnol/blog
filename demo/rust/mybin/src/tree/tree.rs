pub(crate) trait Treer: std::fmt::Debug {
    fn find(self, k: String) -> String;
    fn insert(&mut self, k: String, v: String);
}

pub(crate) fn new_tree() -> impl Treer {
    Tree {
        root: Box::new(Node {
            key: "1".to_string(),
            value: "2".to_string(),
        }),
    }
}

#[derive(Debug)]
struct Tree {
    root: Box<Node>, // 使用Box，避免生命周期标注
}

#[derive(Debug)]
struct Node {
    key: String,
    value: String,
}

impl Treer for Tree {
    fn find(self, k: String) -> String {
        if self.root.key == k {
            self.root.value
        } else {
            "".to_string()
        }
    }
    fn insert(&mut self, k: String, v: String) {
        self.root.key = k;
        self.root.value = v;
    }
}

#[cfg(test)]
mod tests {
    use crate::{new_tree, tree::tree::Treer};

    #[test]
    fn tree() {
        let tree = new_tree();
        let value = tree.find("1".to_string());
        assert_eq!(value, "2".to_string());

        let mut tree1 = new_tree();
        tree1.insert("1".to_string(), "4".to_string());
        let value = tree1.find("1".to_string());
        assert_eq!(value, "4".to_string());
    }
}

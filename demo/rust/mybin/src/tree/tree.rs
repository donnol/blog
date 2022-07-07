
pub trait Treer<'a> : std::fmt::Debug{
    fn find(&self, k :&str) -> &str;
    fn insert(&self, k:& str, v :&str);
}

pub(crate) fn new_tree<'a>() -> impl Treer<'a> {
    Tree{root: &Node { key: "1", value: "2" }}
}

#[derive(Debug)]
struct Tree<'a> {
    root :&'a Node<'a>
}

#[derive(Debug)]
struct Node<'a> {
    key : &'a str,
    value : &'a str
}

impl Treer<'a> for Tree<'a> {
    fn find(&self, k :&str) -> &str{
        if self.root.key == k {
            self.root.value
        }else{
            ""
        }
    }
    fn insert(&self, k:&str, v :&str) {
        self.root.key = k;
        self.root.value = v;
    }
}

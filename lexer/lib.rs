fn main() {
    let s = r#"fn main() {
    let s = r#"{code}"#;
    println!("{}", s.replace("{code}", s));
}"#;
    println!("{}", s.replace("{code}", s));
}
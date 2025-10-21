use std::time::{SystemTime, UNIX_EPOCH};
fn main(){
  let s = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_secs();
  println!("Project: ${REPO_NAME}");
  println!("Epoch seconds: {}", s);
}

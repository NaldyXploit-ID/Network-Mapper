export function demo(){
  const ts = new Date().toISOString();
  return { ts, hash: ts.split('').reduce((a,c)=>a + c.charCodeAt(0),0).toString(16) };
}
if (require.main === module) console.log(demo());

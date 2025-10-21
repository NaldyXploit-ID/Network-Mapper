(function(){
  const btn=document.getElementById('run'), out=document.getElementById('out');
  function demo(){
    const ts = new Date().toISOString();
    out.textContent = "demo at " + ts;
  }
  if(btn) btn.addEventListener('click', demo);
})();

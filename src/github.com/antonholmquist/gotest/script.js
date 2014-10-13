

(function(base64Content) {
 

  var content = atob(base64Content);

  var scriptTag = document.currentScript;

  scriptTag.parentNode.innerHTML = content;

})("<base_64_content>");


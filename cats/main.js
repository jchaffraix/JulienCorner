function $(id) {
  return document.getElementById(id);
}
Handlebars.registerHelper('count', function (aVec) {
    return aVec.length;
});

function shouldIgnorePhotosBasedOnTag(tags) {
  // Ignore second life renderings (there is a lot of those...).
  if (tags.indexOf("secondlife") != -1 || tags.indexOf("second life") != -1 || tags.indexOf("SL") != -1) {
    return true
  }

  if (tags.indexOf("anime") != -1) {
    return true
  }

  if (tags.indexOf("sexy") != -1 || tags.indexOf("girls") != -1) {
    return true
  }

  return false;
}

function filterPhotos(photos) {
  return photos.filter((element) => { return !shouldIgnorePhotosBasedOnTag(element.tags); });
}

function loadImages() {
  const url = "https://www.flickr.com/services/rest/?method=flickr.photos.search&api_key=bbf45c14f25fb5f5b1ee6da824698e04&tags=sphynx%2Chairless%2Cchickencat&text=cat&media=photos&extras=tags%2Cmachine_tags&format=json&nojsoncallback=1";
  fetch(url)
    .then((response) => response.json())
    .then((result) => {
      const template = Handlebars.compile($("image-template").innerHTML);
      if (result.stat != "ok") {
        $("images").innerHTML = "Error from the Flickr API (" + result.toString() + ")";
        return;
      }

      // TODO: For debugging.
      window.result = result;
      result.photos.photo = filterPhotos(result.photos.photo)
      const rendered = template(result.photos);
      $("images").innerHTML = rendered;
    }).catch((error) => {
      $("images").innerHTML = "Error loading from Flickr: " + error;
    });
}

window.addEventListener("load", loadImages);

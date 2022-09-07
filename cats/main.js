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

function shouldIgnoreAuthor(author) {
  // Spammy or just not the right style.
  return author == "194326640@N04" || author == "183799697@N04" || author == "166977808@N03";
}

function filterPhotos(photos) {
  return photos.filter((photo) => {
    if (shouldIgnorePhotosBasedOnTag(photo.tags)) {
      return false;
    }

    if (shouldIgnoreAuthor(photo.owner)) {
      return false;
    }

    return true;
  });
}

function loadImages() {
  const params = new URLSearchParams(window.location.search);
  const page = params.get("page") || 1;
  const per_page = params.get("per_page") || 100;

  const url = "https://www.flickr.com/services/rest/?method=flickr.photos.search&api_key=bbf45c14f25fb5f5b1ee6da824698e04&tags=sphynx%2Chairless%2Cchickencat&text=cat&media=photos&extras=tags%2Cmachine_tags&page=" + page + "&per_page=" + per_page + "&format=json&nojsoncallback=1";
  fetch(url)
    .then((response) => response.json())
    .then((result) => {
      const template = Handlebars.compile($("image-template").innerHTML);
      if (result.stat != "ok") {
        $("images").innerHTML = "Error from the Flickr API (" + result.toString() + ")";
        return;
      }

      let photos = result.photos;
      photos.photo = filterPhotos(photos.photo)
      // Fill prev/next here to simplify our logic.
      if (photos.page > 1) {
        photos.previous = photos.page - 1;
      }
      if (photos.page < photos.pages) {
        photos.next = photos.page + 1;
      }
      const rendered = template(photos);
      $("images").innerHTML = rendered;
    }).catch((error) => {
      $("images").innerHTML = "Error loading from Flickr: " + error;
    });
}

window.addEventListener("load", loadImages);

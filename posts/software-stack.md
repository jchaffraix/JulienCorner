# Software Stack

The goal of this website is to explore lightweight options for hosting content.

## Docker
The site is built into a barebone Alpine Linux docker image.

The image contains a custom Go binary to server static pages.

The image is currently hosted on Google Cloud, but the image should be self-contained and could be deployed anywhere.

## Markdown

Pages and posts are written in Markdown. This is generated as part of building the image. This step is handled through a simple shell script that pulls the dates from Git.

## Source
You can find the [code on GitHub](https://github.com/jchaffraix/JulienCorner.git).

The directory have the following meaning:
* `html`: any static resources to be copied as-is into the image.
* `pages`: Markdown pages
* `posts`: Markdown blog posts
* `src`: Go code to serve the website
* `tools`: any tools used for building or deploying the image

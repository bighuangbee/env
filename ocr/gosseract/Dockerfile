FROM golang:1.17.3

RUN apt-get update -qq

RUN apt-get install -y -qq libtesseract-dev libleptonica-dev


ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/4.00/tessdata/

# Load languages.
# See https://github.com/tesseract-ocr/tessdata for the list of available languages.
RUN apt-get install -y -qq \
  tesseract-ocr-eng \
  tesseract-ocr-deu \
  tesseract-ocr-chi-sim \
  tesseract-ocr-chi-tra


#检查「tesseract」支持的语言
# tesseract --list-langs

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get github.com/otiai10/gosseract/v2

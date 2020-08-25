//base64 转码和解码
let base64Encode = ""
let encodedData = window.btoa(unescape(encodeURIComponent(data)));
let base64Decode = ""
let decodedData = decodeURIComponent(escape(window.atob(base64Decode)))
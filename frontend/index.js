var shortUrl = "";

const CreateShortUrl = async () => {
    url = document.getElementById("inputUrl").value
    reqBody = {"long_url":url}
    console.log(reqBody)
    const response = await fetch('http://localhost:8080/create-short-url', {
      method: 'POST',
      mode:"cors",
      body:JSON.stringify(reqBody),
      headers: {
        'Content-Type': 'application/json'
      }
    });
    const myJson = await response.json()
    console.log(myJson)
    ShowShortUrl(myJson.short_url)
    ShowLongUrl(url)
    return myJson.short_url
  }

  const ShowShortUrl = (url) => {
    document.getElementById("shortUrl").innerHTML = "<h3> Kısaltılmış Link: "+url+"</h3>"
    shortUrl = url;
  }

  const ShowLongUrl = (url) => {
    document.getElementById("longUrl").innerHTML = "<h3> Uzun Link: "+url+"</h3>"
  }

const RedirectMainUrl = async () => {
  const response = await fetch(shortUrl,{
    method: 'GET',
    mode:"no-cors",
    headers: {
      'Content-Type': 'application/json',
    }
  })
  .then(response =>response.json())
  .then(data =>console.log(data))
  document.getElementById("buton").addEventListener("click",function() {window.open(response.long_url, '_blank').focus();});
}
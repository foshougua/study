let myHeading = document.querySelector('h1');
console.log(myHeading.textContent);
myHeading.textContent = 'Hello world!';

let myImg = document.querySelector("img");
myImg.onclick = function(){
    let mySrc = myImg.getAttribute("src");
    console.log(mySrc);
    myImg.setAttribute('src','../image/dongman.jpg');
}

if(!localStorage.getItem('name')) {
    setUserName();
  } else {
    let storedName = localStorage.getItem('name');
    myHeading.textContent = '您好：' + storedName;
}

let myButton = document.querySelector("button");

function setUserName() {
    let myName = prompt('请输入你的名字: ');
    
    localStorage.setItem('name',myName);
    myHeading.textContent = '您好：' + myName;
}
myButton.onclick = function() {
    setUserName()
}

loginBtn = document.getElementById('login_box').getElementsByTagName('div').item(2).querySelector('button')
loginBtn.onclick = function(){
    $.ajax({
            url: "http://127.0.0.1:8080/ping",
            data: {name: 'jenny'},
            type: "GET",
            dataType: "json",
            success: function(data) {
                console.log(data.name)
                alert(data.name)
                window.location.href = "../html/account.html"
            },
            error: function(e) {
                console.log('error',e)
            }
        }
    )
}





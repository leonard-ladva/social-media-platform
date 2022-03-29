startListeners();

function startListeners() {
    postlike = document.getElementsByClassName("iconbutton")[0]
    postdislike = document.getElementsByClassName("iconbutton")[1]
    postlikelistener();
    postdislikelistener();
    comments = document.getElementsByClassName("commentlikebox");
    postid = document.getElementsByTagName("input")[1].value
    commentListeners();
}

//post likes
function postlikelistener() {
    if (postlike) {
        postlike.addEventListener("click", function(e) {
            e.preventDefault();
            postreaction = document.getElementsByTagName("input")[0].value
            formData = { "commentreaction": postreaction }
            fetch('/post?id=' + postid, {
                method: 'POST', // or 'PUT'
                body: new URLSearchParams("postreaction=" + postreaction),
            }).then(response => response.text()).then(body => {
                document.body.innerHTML = body
                startListeners();
            });
        });
    }
}

function postdislikelistener() {
    if (postdislike) {
        postdislike.addEventListener("click", function(e) {
            e.preventDefault();
            postreaction = document.getElementsByTagName("input")[2].value
            formData = { "commentreaction": postreaction }
            fetch('/post?id=' + postid, {
                method: 'POST', // or 'PUT'
                body: new URLSearchParams("postreaction=" + postreaction),
            }).then(response => response.text()).then(body => {
                document.body.innerHTML = body
                startListeners();
            });
        });
    }
}


//comment likes
function commentListeners() {
    Array.from(comments).forEach(element => {

        like = element.getElementsByClassName("iconbutton")[0]
        dislike = element.getElementsByClassName("iconbutton")[1]

        //function for posting like
        dislike.addEventListener("click", function(e) {
            commentreaction = element.getElementsByTagName("input")[2].value
            commentid = element.getElementsByTagName("input")[1].value
            e.preventDefault();
            let formData = new FormData();
            formData = { "commentreaction": commentreaction, "commentid": commentid }
            fetch('/post?id=' + postid, {
                method: 'POST', // or 'PUT'
                body: new URLSearchParams("commentreaction=" + commentreaction + "&commentid=" + commentid),
            }).then(response => response.text()).then(body => {
                document.body.innerHTML = body
                startListeners();
            });
        })

        like.addEventListener("click", function(e) {
            commentreaction = element.getElementsByTagName("input")[0].value
            commentid = element.getElementsByTagName("input")[1].value
            e.preventDefault();
            let formData = new FormData();
            formData = { "commentreaction": commentreaction, "commentid": commentid }
            fetch('/post?id=' + postid, {
                method: 'POST', // or 'PUT'
                body: new URLSearchParams("commentreaction=" + commentreaction + "&commentid=" + commentid),
            }).then(response => response.text()).then(body => {
                document.body.innerHTML = body
                startListeners();
            });
        })
    });
}
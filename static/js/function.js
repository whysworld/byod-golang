function setContentHeight() {
    var offset = 80;
    var $containerHeight = $(".container").outerHeight();
    var $welcomeWrapperHeight = $(".welcome-wrapper").outerHeight();
    var $titleHeight = $(".content-title").outerHeight();
    var $contentHeight = $containerHeight - $welcomeWrapperHeight - $titleHeight - offset
    $(".content").height($contentHeight);
}
$(document).ready(function () {
    setContentHeight();
    $("#registration-login-btn").on("click", function(e){
        e.preventDefault();
        //http://guestportal/registration/info/ -----> http://guestportal/registration/info
        var pathName = window.location.pathname.replace(/\/\/*/g,"/").replace(/\/+$/,"")
        //http://guestportal/registration/info/ -----> info
        var search = pathName.split("/").slice(-1).toString()
        var url = window.location.href.replace(search, "login")
        window.location.href = url
    })
    $("#registration-signup-btn").on("click", function(e){
        e.preventDefault();
        //http://guestportal/registration/login/ -----> http://guestportal/registration/login
        var pathName = window.location.pathname.replace(/\/\/*/g,"/").replace(/\/+$/,"")
        //http://guestportal/registration/info/ -----> info
        var search = pathName.split("/").slice(-1).toString()
        var url = window.location.href.replace(search, "info")
        window.location.href = url 
    })

    $(".carousel-control-prev").on("click", function(e) { 
        $("#usersCarousel").carousel('prev') 
    })
    $(".carousel-control-next").on("click", function(e) { 
        $("#usersCarousel").carousel('next') 
    })
    $(window).resize(function () {
        setContentHeight();
    })
});
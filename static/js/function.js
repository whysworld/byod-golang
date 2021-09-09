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
    $("#hotspot-accept").on("click", function(e){
        e.preventDefault();
        window.location.href = "/hotspot/accept"
    })
    $("#hotspot-back").on("click", function(e){
        e.preventDefault();
        window.location.href = "/hotspot/home"
    })
    $("#hotspot-decline").on("click", function(e){
        e.preventDefault();
        window.location.href = "/hotspot/decline"
    })
    $("#registration-accept").on("click", function(e){
        e.preventDefault();
        window.location.href = "/registration/info"
    })
    $("#registration-form").on("submit", function(e){
        e.preventDefault();
        window.location.href = "/registration/login"
    })
    $("#registration-login").on("submit", function(e){
        e.preventDefault();
        window.location.href = "/registration/accept"
    })
    $("#sponsor-login").on("submit", function(e){
        e.preventDefault();
        window.location.href = "/sponsor/users"
    })
    $("#sponsor-cancel").on("click", function(e){
        e.preventDefault();
        window.location.href = "/sponsor/users"
    })
    $(window).resize(function () {
        setContentHeight();
    })
});
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
    $("#logout").on("click", function(e){
        e.preventDefault();
        window.location.href = "/sponsor/logout"
    })
    $(window).resize(function () {
        setContentHeight();
    })
});
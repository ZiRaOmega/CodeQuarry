input = document.getElementById('search-bar-input');

$(document).ready(function () {
    console.log("ready!");
    $("#search-bar-input").on("input", function () {
        const search = input.value;
        console.log(search);
        if (search.length > 0) {
            $.ajax({
                url: "/search",
                method: "POST",
                data: { search },
                success: function (data) {
                    $("#search-results").html(data);
                },
                error: function (error) {
                    console.error(error);
                },
            });
        } else {
            $("#search-results").html("");
        }
    });
}
);
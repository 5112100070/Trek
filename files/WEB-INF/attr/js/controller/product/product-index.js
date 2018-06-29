$(document).ready(function(){
    // Add scrollspy to <body>
  $('body').scrollspy({target: ".navbar", offset: 50});   

  // Add smooth scrolling on all links inside the navbar
  $("#myNavbar a").on('click', function(event) {
    // Make sure this.hash has a value before overriding default behavior
    if (this.hash !== "") {
      // Prevent default anchor click behavior
      event.preventDefault();

      // Store hash
      var hash = this.hash;

      // Using jQuery's animate() method to add smooth page scroll
      // The optional number (800) specifies the number of milliseconds it takes to scroll to the specified area
      $('html, body').animate({
        scrollTop: $(hash).offset().top
      }, 800, function(){
   
        // Add hash (#) to URL when done scrolling (default click behavior)
        window.location.hash = hash;
      });
    }  // End if
  });

  loadlistProduct();
});

function loadlistProduct(){
    promise = getProductForMP();

    promise.done(function(response){
        if(response.data.length == 0){
            return;
        }

        for(var i=response.data.length-1;i>=0;i--){
            var template_choice = ``+
                `<div class="col-lg-3">` + 
                    `<div class="features-icons-item mx-auto">` + 
                        `<div class="features-icons-icon d-flex">` +
                            `<img src="img/dimanasaja.png" style="min-width:10rem;"alt="kapan saja dimana saja">` +
                        `</div>` + 
                        `<h5 style="padding-top:1rem">Bor Listrik</h3>` + 
                        `<p class="desc" style="text-align:center;">`+ response.data[i].price_to_sell +`/Minggu` +
                    `</div>` +
                `</div>`;

            $("#parent-list-products").after($(template_choice));
        }
    });
    
}
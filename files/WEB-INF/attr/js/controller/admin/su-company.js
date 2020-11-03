// only role client you can set member of company
$("#role").change(function(){
    if($("#role").val() == 2) {
        $("#init-member-of-company").attr("selected", true);

        $("#form-group-member-of-company").attr('style', 'display: block');
    } else {
        $("#init-member-of-company").attr("selected", true);

        $("#form-group-member-of-company").attr('style', 'display: none');
    }
});

function FetchCompanyByAdminRole(){
    StartLoading();

        var params = new URLSearchParams({
            page: 1,
            rows: 50,
            role: 1
        })
        var url = product_url + '/admin/v1/get-list-company?' + params;
        

        var promise = $.ajax({
            url: url,
            type: 'GET',
            crossDomain: true,
            xhrFields: {
                withCredentials: true
            },
            headers: {
                "Authorization": GetCookie('_CGX_DEV_'),
                "Accept": "application/json"
            }
        });

    promise.done(function(response){
        response = response.data.companies;
        response.forEach(element => {                    
            // generate new component for modal input
            var comp = `
                <option value="` + element.company_id + `">` + element.company_name + `</option>
            `;
            
            // create and push component for modal
            $("#init-member-of-company").after(comp);
        });
        if(response.error!=null){
            setError(true, response.error.detail);
        }
        FinishLoading();
    }).fail(function(response){
        setError(true, defaultServerError);
        FinishLoading();
    });

    return
}
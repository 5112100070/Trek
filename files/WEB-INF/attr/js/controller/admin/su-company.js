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
            role: 1,
            is_enabled: "1"
        })
        var path = '/admin/v1/get-list-company';
        var url = product_url + path + '?' + params;
        
        let now = moment();
        // 02 Jan 06 15:04 MST
        let stdHeaderTime = now.format("DD MMM gg HH:mm");
        // datetime formatted version
        let stdFormatTime = moment(stdHeaderTime);
        let hash = generateHMACHash("GET", path, stdFormatTime.format("X") , "");

        var promise = $.ajax({
            url: url,
            type: 'GET',
            crossDomain: true,
            xhrFields: {
                withCredentials: true
            },
            headers: {
                "Authorization": GetSessionBasedOnEnv(),
                "Accept": "application/json",
                "User-Agent-2": "cgx",
                "Authorization-2": hash,
                "Date-Auth": stdHeaderTime+" WIB"
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
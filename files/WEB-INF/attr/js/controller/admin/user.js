function loadlistUser(totalRequest = 8){
    promise = getUserList(totalRequest, "DESC");

    promise.done(function(response){
        if(response.data.length == 0){
            return;
        }

        var typeUser;
        for(var i=response.data.length-1;i>=0;i--){
            if(response.data[i].type == 0){
                typeUser = "admin Trek";
            } else if(response.data[i].type == 1){
                typeUser = "user biasa";
            } else {
                typeUser = "tak terindentifikasi";
            }

            var template_choice = ``+
                `<tr style="cursor:pointer;" onClick="javascript:GoToIndex('admin/user/edit?user-id=` + response.data[i].user_id + `')">` + 
                    `<td>` + (i+1) + `</td>` +
                    `<td>` + response.data[i].user_id+ `</td>` +
                    `<td>` + response.data[i].fullname + `</td>` +
                    `<td>` + response.data[i].username + `</td>` +
                    `<td>` + response.data[i].status + `</td>` +
                    `<td>` + typeUser + `</td>` +
                    `<td>` + response.data[i].create_time + `</td>` +
                `</tr>`;

            $("#parent-list-user").after($(template_choice));
        }
    });
}

function loadUserDetailByID(userID){
    promise = getUserByID(userID);

    promise.done(function(response){
        $("#fullname").val(response.data.fullname);
        $("#username").val(response.data.username);
        $("#status").val(response.data.status);
        $("#type").val(response.data.type);
        
        $("#create-time").val(response.data.create_time);
        $("#update-time").val(response.data.update_time);
        $("#user-img").attr("src", base_url + response.data.img_url);
    })
}
const type_success = "success";
const type_error = "error";
const type_question = "question";
const type_warning = "warning";

function SimpleSwalWithoutDesc(title, type, confirmButton){
    return swal({
        title: title,
        type: type,
        confirmButtonText: confirmButton,
        customClass: "animated tada",
        animation: false,
        imageUrl: '/img/logo-cgx.png',
        imageWidth: 150,
        imageHeight: 50,
        background: 'rgb(255, 204, 0)',
        imageAlt: 'Trek Logo',
    });
}

function SimpleSwal(title, text, type, confirmButton){
    return swal({
        title: title,
        text: text,
        type: type,
        confirmButtonText: confirmButton,
        animation: false,
        background: 'rgb(255, 204, 0)',
        imageAlt: 'Trek Logo',
    });
}

function SwalConfimationProcess(title, question, type, confirmButton, cancelButton){
    return swal({
        title: title,
        text: question,
        type: type,
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: confirmButton,
        cancelButtonText: cancelButton,
        customClass: "animated tada",
        animation: false,
        imageUrl: '/img/logo-cgx.png',
        imageWidth: 150,
        imageHeight: 50,
        background: 'rgb(255, 204, 0)',
        imageAlt: 'Trek Logo',
      });
}

function validateType(type){
    if(type==type_success || type==type_error || type==type_question || type == type_warning){
        return type;
    } else {
        return type_success;
    }
}
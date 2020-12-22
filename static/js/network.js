function httpPost(URL, PARAMS) {
    var temp = document.createElement("form");
    temp.action = URL;
    temp.method = "post";
    temp.style.display = "none";

    for (var x in PARAMS) {
        var opt = document.createElement("textarea");
        opt.name = x;
        opt.value = PARAMS[x];
        temp.appendChild(opt);
    }

    document.body.appendChild(temp);
    temp.submit();

    return temp;
}

function toMonitorCompanyList(groupName,riskLevel,updateStatus,province,eventLevel,reportId,riskStatus){

    var params = {
        "groupName": encodeURIComponent(groupName),
        "riskLevel": riskLevel,
        "updateStatus": updateStatus,
        "province": encodeURIComponent(province),
        "eventLevel": eventLevel,
        "reportId": reportId,
        "riskStatus":riskStatus
    };

    httpPost(ctx + "/js/network", params);
}
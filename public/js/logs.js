$(document).ready(function() {
	Calendar.setup({
		weekNumbers : true,
		inputField : "start_time",
		trigger : "start_time",
		dateFormat : "%Y-%m-%d %H:%M:%S",
		showTime : true,
		minuteStep : 1,
		onSelect : function() {
			this.hide();
		}
	});

	Calendar.setup({
		weekNumbers : true,
		inputField : "end_time",
		trigger : "end_time",
		dateFormat : "%Y-%m-%d 23:59:59",
		showTime : true,
		minuteStep : 1,
		onSelect : function() {
			this.hide();
		}
	});
});

//清理日志
function delAll() {
	art.dialog.confirm('你确定要清理日志吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Logs/delAll/",
			data : "",
			success : function(html) {
				var tmp = jQuery.parseJSON(html);
				if (tmp.rtn_code == 0) {
					notice_tips("删除清理日志成功!");
					right_refresh();
				} else {
					notice_tips(tmp.content);
				}
			}
		});
	}, function() {
		notice_tips("你取消了清理日志操作!");
	});
}

//搜索
function Search() {
	url = '/Logs/index/';
	url += 'username:' + $('#username').val() + '/';
	url += 'realname:' + $('#realname').val() + '/';
	url += 'start_time:' + $('#start_time').val() + '/';
	url += 'end_time:' + $('#end_time').val() + '/';

	window.location.href = url;
};
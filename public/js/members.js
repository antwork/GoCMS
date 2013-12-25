/**
 * 用户管理
 */

/**
 * 编辑用户
 */
function edit(uid) {
	if (uid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/Members/edit/' + uid + '/', {
		id : 'user_edit',
		title : '编辑用户',
		width : 700,
		height : 500,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var status = iframe.$('#status').val();

			if (status == '' || status == 'undefined') {
				art.dialog.alert('参数错误!');
				return false;
			}

			var par = [];
			var pars = "uid=" + uid;
			par.push(pars);
			var pars = "status=" + status;
			par.push(pars);
			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Members/edit/",
				data : pars,
				success : function(html) {
					var tmp = jQuery.parseJSON(html);
					if (tmp.rtn_code == 0) {
						window.location.reload();
						notice_tips("编辑用户成功!");
					} else {
						notice_tips(tmp.content);
					}
				}
			});
		},
		cancel : true
	});
}

/**
 * 删除用户
 */
function del(uid) {
	if (uid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除这个用户吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Members/delete/",
			data : "uid=" + uid,
			success : function(html) {
				var tmp = jQuery.parseJSON(html);
				if (tmp.rtn_code == 0) {
					window.location.reload();
					notice_tips("删除成功!");
				} else {
					notice_tips(tmp.content);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除用户操作!");
	});
}
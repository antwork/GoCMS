/**
 * 管理员管理
 */

/**
 * 编辑管理员
 */
function edit(uid) {
	if (uid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/Administrator/edit/' + uid + '/', {
		id : 'admin_edit',
		title : '编辑管理员',
		width : 500,
		height : 400,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var uid = iframe.$('#uid').val();
			var password = iframe.$('#password').val(); // 密码
			var pwdconfirm = iframe.$('#pwdconfirm').val(); // 确认密码
			var email = iframe.$('#email').val(); // E-mail
			var realname = iframe.$('#realname').val(); // 真实姓名
			var roleid = iframe.$('#roleid').val(); // 所属角色

			if (uid == '' || uid == 'undefined') {
				art.dialog.alert('参数错误!');
				return false;
			}

			if (password != pwdconfirm) {
				art.dialog.alert('两次输入的密码不一样!');
				return false;
			}
			if (email == '') {
				art.dialog.alert('请输入邮箱!');
				return false;
			}
			if (realname == '') {
				art.dialog.alert('请输入真实姓名!');
				return false;
			}

			var par = [];
			var pars = "uid=" + uid;
			par.push(pars);
			if (password != '' && pwdconfirm != '') {
				pars = "password=" + password;
				par.push(pars);
				pars = "pwdconfirm=" + pwdconfirm;
				par.push(pars);
			}

			pars = "email=" + email;
			par.push(pars);
			pars = "realname=" + realname;
			par.push(pars);
			pars = "roleid=" + roleid;
			par.push(pars);
			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Administrator/edit/",
				data : pars,
				success : function(html) {
					var tmp = jQuery.parseJSON(html);
					if (tmp.rtn_code == 0) {
						window.location.reload();
						notice_tips("编辑管理员成功!");
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
 * 删除管理员
 */
function delete_admin(uid) {
	if (uid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除这个管理员吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Administrator/delete/",
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
		notice_tips("你取消了删除管理员操作!");
	});
}
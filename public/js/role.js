/**
 * 角色管理
 */

/**
 * 权限设置
 */
function setting_role(roleid, rolename) {
	art.dialog.open('/Role/priv_setting/' + roleid + '/', {
		id : 'setting_role',
		title : '设置《' + rolename + '》',
		width : 700,
		height : 500,
		lock : true
	});
}

/**
 * 表单提交
 * 
 * @returns {Boolean}
 */
function form_submit() {
	var rolename = $.trim($("#rolename").val());
	if (rolename == '') {
		$("#rolename").focus();
		notice_tips("请输入角色名称!");
		return false;
	}

	var description = $.trim($("#description").val());
	if (description == '') {
		$("#description").focus();
		notice_tips("请输入角色描述!");
		return false;
	}
	return true;
}

/**
 * 栏目权限
 * 
 * @param roleid
 * @param rolename
 */
function setting_cat_priv(roleid, rolename) {
	art.dialog.open('/Role/setting_cat_priv/' + roleid + '/', {
		id : 'setting_cat_priv',
		title : '栏目权限《' + rolename + '》',
		width : 700,
		height : 500,
		lock : true
	});
}

/**
 * 删除角色
 * 
 * @param roleid
 */
function delete_role(roleid) {
	if (roleid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除这个角色吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Role/delete/",
			data : "roleid=" + roleid,
			success : function(html) {
				var tmp = jQuery.parseJSON(html);
				if (tmp.rtn_code == 0) {
					notice_tips("删除成功!");
					window.location.reload();
				} else {
					notice_tips(tmp.content);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除角色操作!");
	});
}

/**
 * 设置状态
 * 
 * @param roleid
 * @param disabled
 */
function set_disabled(roleid, disabled) {
	if (roleid == '') {
		notice_tips("参数错误!");
		return false;
	}

	if (disabled == 0) {
		var message = '你确定要启用这个角色及用户吗?';
	} else {
		var message = '你确定要禁用这个角色及用户吗?';
	}

	art.dialog.confirm(message, function() {
		$.ajax({
			type : "POST",
			url : "/Role/setStatus/",
			data : "roleid=" + roleid + "&disabled=" + disabled,
			success : function(html) {
				var tmp = jQuery.parseJSON(html);
				if (tmp.rtn_code == 0) {
					notice_tips("设置成功!");
					window.location.reload();
				} else {
					notice_tips(tmp.content);
				}
			}
		});
	}, function() {
		notice_tips("你取消了设置状态操作!");
	});
}
/**
 * 修改密码
 */

/**
 * 提交检测
 */
function form_submit() {
	var realname = $.trim($("#realname").val());
	if (realname == '') {
		$("#realname").focus();
		notice_tips("请输入真实姓名!");
		return false;
	}

	var email = $.trim($("#email").val());
	if (email == '') {
		$("#email").focus();
		notice_tips("请输入电子邮件!");
		return false;
	}

	var lang = $.trim($("#lang").val());
	if (lang == '') {
		$("#lang").focus();
		notice_tips("请选择语言!");
		return false;
	}

	return true;
}
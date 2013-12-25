/**
 * 管理栏目
 */

// Tab切换
function SwapTab(name, cls_show, cls_hide, cnt, cur) {
	for ( var i = 1; i <= cnt; i++) {
		if (i == cur) {
			$('#div_' + name + '_' + i).show();
			$('#tab_' + name + '_' + i).attr('class', cls_show);
		} else {
			$('#div_' + name + '_' + i).hide();
			$('#tab_' + name + '_' + i).attr('class', cls_hide);
		}
	}
}

/**
 * 提交检测
 */
function form_submit() {
	var catname = $.trim($("#catname").val());
	if (catname == '') {
		$("#catname").focus();
		notice_tips("请输入栏目名称!");
		return false;
	}

	var controller = $.trim($("#controller").val());
	if (controller == '') {
		$("#controller").focus();
		notice_tips("请输入栏目地址!");
		return false;
	}
	return true;
}

/**
 * 删除菜单
 * 
 * @param id
 */
function delete_cate(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Content/delCate/",
			data : "id=" + id,
			success : function(html) {
				var tmp = jQuery.parseJSON(html);
				if (tmp.rtn_code == 0) {
					notice_tips("删除栏目成功!");
					right_refresh();
				} else {
					notice_tips(tmp.content);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除栏目操作!");
	});
}
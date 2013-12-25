/**
 * 公告管理
 */
/**
 * 编辑公告分类
 */
function cate_edit(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/Notice/editCate/' + id + '/', {
		id : 'cate_edit',
		title : '编辑分类',
		width : 300,
		height : 120,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var id = iframe.$('#id').val();
			var name = iframe.$('#name').val(); // 分类名称

			if (id == '' || id == 'undefined') {
				art.dialog.alert('参数错误!');
				return false;
			}

			if (name == '') {
				art.dialog.alert('请输入分类名称!');
				return false;
			}

			var par = [];
			var pars = "id=" + id;
			par.push(pars);
			pars = "name=" + name;
			par.push(pars);
			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Notice/editCate/",
				data : pars,
				success : function(html) {
					var tmp = jQuery.parseJSON(html);
					if (tmp.rtn_code == 0) {
						window.location.reload();
						notice_tips("编辑分类成功!");
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
 * 删除公告分类
 */
function cate_delete(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Notice/deleteCate/",
			data : "id=" + id,
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
		notice_tips("你取消了删除分类操作!");
	});
}

/**
 * 删除公告
 */
function delete_notice(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Notice/delete/",
			data : "id=" + id,
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
		notice_tips("你取消了删除公告操作!");
	});
}
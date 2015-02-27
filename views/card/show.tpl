<ul class="nav nav-tabs">
  <li role="presentation" class="active"><a href="#">分配叮当卡</a></li>
</ul>

<div class="container" id="search_box">
	<form class="form-horizontal" action="{{urlfor "CardController.Search"}}" method="post" name="search_form" id="search_form">
	  <div class="form-group">
	    <div class="col-sm-10">
	    	<div class="control-group">
	      		<input type="text" class="form-control validate" name="search_condition" id="search_condition" data-rules="required"  placeholder="家庭成员姓名/成员手机号/家庭名称" data-display="搜索条件">
	      		<p class="help-block"></p>
			</div>
	    </div>
	    <input type="submit" id="search" class="col-sm-2 btn btn-primary" value="搜索">
	  </div>
	</form>
</div>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">
<html>
<frameset rows="8%,92%">

    <frame src="{{urlfor "MainController.TopBody"}}" scrolling="no" noresize="noresize" name="top"/>

    <frameset cols="120px,*">
        <frame src="{{urlfor "MainController.LeftBody"}}" noresize="noresize"  name="left"/>
        <frame src="{{urlfor "MainController.RightBody"}}" name="right"/>
    </frameset>

</frameset>
<script type="application/javascript">

</script>
</html>
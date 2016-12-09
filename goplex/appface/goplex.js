$(document).ready(function() {

        // Get the screen resolution
        var width = $(window).width();
        var height = $(window).height();
        var current_url = window.location.href;
        current_url = current_url.split('_');
        if (current_url[1] != 'r') {
            var wh = '_r_'+width+'_x_'+height;
            window.location.href = window.location.href+wh;
            console.log(wh);
        }
        // Hide the cursor rect
        $('#cursorRect').hide();
        // Event listeners
        $('img').on('click', function(event) {
            // Coordinates convertion
            // Extract the current coordinates from the url
            var url = window.location.href;
            var coordinates = url.split("_");
            var xresolution = parseFloat(coordinates[2]);
            var yresolution = parseFloat(coordinates[4]);
            var x_min = parseFloat(coordinates[6]);
            var x_max = parseFloat(coordinates[8]);
            var y_min = parseFloat(coordinates[10]);
            var y_max = parseFloat(coordinates[12]);
            var new_x = event.pageX - this.offsetLeft;
            var new_y = event.pageY - this.offsetTop;
            x_mn = x_min+((new_x-80)/xresolution)*(x_max-x_min);
            x_mx = x_min+((new_x+80)/xresolution)*(x_max-x_min);
            y_mn = y_min+((new_y-80)/yresolution)*(y_max-y_min);
            y_mx = y_min+((new_y+80)/yresolution)*(y_max-y_min);
            var url_param = "_"+String(x_mn.toFixed(6))+"_"+String(x_mx.toFixed(6))+"_"+String(y_mn.toFixed(6))+"_"+String(y_mx.toFixed(6));
            window.location.href = window.location.href+url_param;
        });
        $('img').on('mousemove', function(event) {
            $(this).css('cursor', 'none');
            var x = event.pageX - 80;
            var y = event.pageY - 80;
            $('#cursorRect').css('top', y).css('left', x);
            // Coordinates convertion
            // Extract the current coordinates from the url
            var url = window.location.href;
            var coordinates = url.split("_");
            var xresolution = parseFloat(coordinates[2]);
            var yresolution = parseFloat(coordinates[4]);
            var x_min = parseFloat(coordinates[6]);
            var x_max = parseFloat(coordinates[8]);
            var y_min = parseFloat(coordinates[10]);
            var y_max = parseFloat(coordinates[12]);
            var new_x = event.pageX - this.offsetLeft;
            var new_y = event.pageY - this.offsetTop;
            x_mn = x_min+((new_x-80)/xresolution)*(x_max-x_min);
            x_mx = x_min+((new_x+80)/xresolution)*(x_max-x_min);
            y_mn = y_min+((new_y-80)/yresolution)*(y_max-y_min);
            y_mx = y_min+((new_y+80)/yresolution)*(y_max-y_min);
            $("input[name$='Xmin']").val(x_mn.toFixed(6));
            $("input[name$='Ymin']").val(y_mn.toFixed(6));
            $("input[name$='Xmax']").val(x_mx.toFixed(6));
            $("input[name$='Ymax']").val(y_mx.toFixed(6));
        });
        $('img').on('mouseenter', function(event) {
            $('#cursorRect').show();
        });
        $('img').on('mouseleave', function(event) {
            $('#cursorRect').hide();
            // Extract the current coordinates back from the url
            var url = window.location.href;
            var coordinates = url.split("_");
            var x_min = parseFloat(coordinates[6]);
            var x_max = parseFloat(coordinates[8]);
            var y_min = parseFloat(coordinates[10]);
            var y_max = parseFloat(coordinates[12]);
            $("input[name$='Xmin']").val(x_min.toFixed(6));
            $("input[name$='Ymin']").val(y_min.toFixed(6));
            $("input[name$='Xmax']").val(x_max.toFixed(6));
            $("input[name$='Ymax']").val(y_max.toFixed(6));
        });
        // Reload the image
        //$('#set').attr('src', "set/mandelbrot.png" + new Date());
});

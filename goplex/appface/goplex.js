$(document).ready(function() {

        // Get the screen resolution
        var width = $(window).width();
        var height = $(window).height();
        var current_url = window.location.href;
        current_url = current_url.split('_');
        if (current_url.length < 2) {
            var wh = '_r_'+width+'_x_'+height;
            window.location.href = window.location.href+wh;
            //console.log(wh);
        }
        else {
            var background_color = current_url[13];
            var color = current_url[14];
            switch (background_color) {
                case 'b':
                    $('body').css('background', 'black');
                    $('#cursorRect').css('border-color', 'white');
                    break;
                case 'w':
                    $('body').css('background', 'white');
                    $('#cursorRect').css('border-color', 'black');
                    break;
            }
        }
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
            // Hide the cursor rect
            //$('#cursorRect').hide();
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
            var x = event.pageX - 80;
            var y = event.pageY - 80;
            $('#cursorRect').css('top', y).css('left', x);
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
        // Form submit
        $('form').submit(function(event) {
            event.preventDefault();
            current_url[6] = $("input[name$='Xmin']").val();
            current_url[8] = $("input[name$='Xmax']").val();
            current_url[10] = $("input[name$='Ymin']").val();
            current_url[12] = $("input[name$='Ymax']").val();
            var new_url = "";
            for (var i = 0; i < current_url.length; i++) {
                if (i > 0) new_url += '_';
                new_url += current_url[i];
            }
            window.location.href = new_url;
        });
        // Color form
        $('h2').on('click', function() {
            $('#colorForm').toggle();
        });
        $('#colorForm').submit(function(event) {
            event.preventDefault();
            current_url[13] = $('#background').val();
            var new_url = "";
            for (var i = 0; i < current_url.length; i++) {
                if (i > 0) new_url += '_';
                new_url += current_url[i];
            }
            window.location.href = new_url;
        });
});

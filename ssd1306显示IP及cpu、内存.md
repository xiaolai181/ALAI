# ssd1306显示IP及cpu、内存

导入PIL包

``` 
from PIL import Image
from PIL import ImageDraw
from PIL import ImageFont
```

导入adafruit_SSD1306

``` 
import Adafruit_SSD1306
```

time

``` 
import time
```

subprocess执行linux命令

```
import subprocess
```

设置128x32显示器配置，i2c总线设为1：

``` 
# 128x32 display with hardware I2C:
disp = Adafruit_SSD1306.SSD1306_128_32(rst=None, i2c_bus=1, gpio=1) # setting gpio to 1 is hack to avoid platform detection
```

通过i2c连接OLED模块

``` 
while True:
    
    try:
        # Try to connect to the OLED display module via I2C.
        disp.begin()
    except OSError as err:
        print("OS error: {0}".format(err))
        time.sleep(10)
    else:
        break
```

清除显示，通过PIL创建绘图对象

``` 
# Clear display.
disp.clear()
disp.display()

# Create blank image for drawing.
# Make sure to create image with mode '1' for 1-bit color.
width = disp.width
height = disp.height
image = Image.new('1', (width, height))

# Get drawing object to draw on image.
draw = ImageDraw.Draw(image)
# Draw a black filled box to clear the image.
draw.rectangle((0,0,width,height), outline=0, fill=0)

# Draw some shapes.
# First define some constants to allow easy resizing of shapes.
padding = -2
top = padding
bottom = height-padding
# Move left to right keeping track of the current x position for drawing shapes.
x = 0

# Load default font.
font = ImageFont.load_default()
```

读取数据并使用drwa对象输出文本

```
while True:

    # Draw a black filled box to clear the image.
    draw.rectangle((0,0,width,height), outline=0, fill=0)

    # Shell scripts for system monitoring from here : https://unix.stackexchange.com/questions/119126/command-to-display-memory-usage-disk-usage-and-cpu-load
    cmd = "top -bn1 | grep load | awk '{printf \"CPU Load: %.2f\", $(NF-2)}'"
    CPU = subprocess.check_output(cmd, shell = True )
    cmd = "free -m | awk 'NR==2{printf \"Mem: %s/%sMB %.2f%%\", $3,$2,$3*100/$2 }'"
    MemUsage = subprocess.check_output(cmd, shell = True )
    cmd = "df -h | awk '$NF==\"/\"{printf \"Disk: %d/%dGB %s\", $3,$2,$5}'"
    Disk = subprocess.check_output(cmd, shell = True )

    # Write two lines of text.

    draw.text((x, top),       "eth0: " + str(get_ip_address('eth0')),  font=font, fill=255)
    draw.text((x, top+8),     "wlan0: " + str(get_ip_address('wlan0')), font=font, fill=255)
    draw.text((x, top+16),    str(MemUsage.decode('utf-8')),  font=font, fill=255)
    draw.text((x, top+25),    str(Disk.decode('utf-8')),  font=font, fill=255)

    # Display image.
    disp.image(image)
    disp.display()
    time.sleep(1)
```




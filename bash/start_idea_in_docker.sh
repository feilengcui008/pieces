docker run \
    -dti \
    -e DISPLAY=$DISPLAY \
    -v /tmp/.X11-unix:/tmp/.X11-unix \
    -v /home/tan/me/workspace/java:/home/developer/dev \
    --name idea smilingrob/intellij-ce:jdk7-15.0.4

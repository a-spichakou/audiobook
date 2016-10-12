for f in ../output/output*; do ffmpeg -i $f -i ../sounds/track.mp3 -filter_complex 'volume=5[a0];[1:a]volume=1[a1];[a0][a1]amix=inputs=2:duration=first' $f.mp3 ; done


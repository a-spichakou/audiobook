for n in {0..6}; do
    for m in {0..9}; do
      ffmpeg -f concat -i <(for f in ./tts$n$m*.ogg; do echo "file '$PWD/$f'"; done) -c copy output$n$m.ogg;
    done
done
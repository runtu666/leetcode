for file in ./*
do
  if test -d $file
      then
          mv $file/main.go ./$file.go
  fi
done
IMAGEDIR=$PWD/$(dirname $0)
cd $IMAGEDIR
rm -rf *.go ./f2bs
curl -Ls https://raw.githubusercontent.com/hajimehoshi/file2byteslice/master/main.go > f2bs.go
go build f2bs.go
rm f2bs.go
for ext in png jpg jpeg gif; do
	echo $ext\s:
	for file in $(find . -name *.$ext); do
		file=$(basename $file)
		name=$(echo $file | sed "s,.$ext\$,,")_$ext
		var=$(echo $name | sed -r 's,(^|_)([a-z]),\U\2,g')
		pkg=$(dirname $0)
		cmd="./f2bs -input $file -output $name.go -var $var -package $pkg"
		echo $cmd
		$cmd
	done
done
rm ./f2bs

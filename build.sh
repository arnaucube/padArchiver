echo "starting build process"
rm -rf build
mkdir build

cd webServer
echo "building webServer"
go build
cp webServer ../build
cp -r templates ../build
cd ..

cd listPadsImporter
echo "building listPadsImporter"
go build
cp listPadsImporter ../build
cd ..

cd apiServer
echo "building apiServer"
go build
cp apiServer ../build
cd ..


cd cli
echo "building cli"
go build
cp cli ../build
cd ..

echo "build process complete, all the necessary files are in the /build directory"

# Запускает все юнит тесты, что есть в проекте
echo "Start unit tests"
cd pkg
go test -count=1 .
echo "Start bench tests"
go test -bench=. -benchmem
echo "Done"
export SONAR_BASE=./

coverage_float=$(go tool cover -func ${SONAR_BASE}/coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
coverage=${coverage_float%.*}
echo "Current coverage is: $coverage"
threshold=80
if [ "$coverage" -lt $threshold ]; then
  echo "Test coverage threshold of $threshold% was not reached, current coverage is $coverage%"
  exit 1
fi;
if [[ $# -eq 0 ]] ; then
    echo "usage:  ./api_client.sh Rowen kathryn@email.com "heart+surgery" \$100 true" 
    exit 
fi

name=$1
email=$2
procedure=$3
cost=$4
paid=$5

payload="petName=$name&ownerEmail=$email&procedure=$procedure&cost=$cost&isPaid=$paid"
curl -d $payload -X POST http://localhost:3000/api
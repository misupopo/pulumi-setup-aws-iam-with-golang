# 発行されたiamのリストを確認
get-iam-user-list:
	aws iam list-users --query 'Users[].UserName'

get-iam-group-list:
	aws iam list-groups \
    --query 'Groups[].GroupName'

# 指定したawsのuserがどこのグループに所属しているか確認
# get-belongs-group userName=name
get-belongs-group:
	aws iam list-groups-for-user \
    --user-name ${userName} \
    --query 'Groups[].GroupName'

# 実行前にどのresourceがあたるかを確認することができる
# terraform planと同じにあたる
pulumi-preview:
	pulumi preview

# アタッチされているポリシーなどを確認する方法
# https://qiita.com/isobecky74/items/92d35fa1d3063fe64dc4

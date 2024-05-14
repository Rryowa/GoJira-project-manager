## **RESTful** Jira **JWT** API in **Golang**<br>
**No Web frameworks**<br>
**Layered** architecture implementation<br>
>We built an optimised Golang image using Multi-stage builds. This is because using **golang:latest** creates an image that consumes a lot of memory and hence not production friendly. So we used **golang:alpine** which is very light. A whopping **17.9MB** image. Impressive!
>Another thing worth noting is the **.env** file. You can see how we added it to the **Dockerfile**. This is essential for our environmental variables.
>Observe the DB_HOST. This is gotten from the name of the service in the docker-compose.yml file. HUH.
***RUN***<br>
##
<tab><tab>docker-compose build<br>
##
<tab><tab>docker-compose up<br><br>
**Thunder Client:**<br>
POST<br>
##
<tab><tab>http://[::1]:3000/api/v1/users/register<br>
with json body email, password<br>
```json
{
  "email":"baaaa1bbbbb@gmail.com",
  "Password":"1234567890"
}
```
COPY token from response(or cookies)<br>
POST<br>
##
<tab><tab>http://[::1]:3000/api/v1/tasks
<br>
with copied token in Authorization header ^_^<br>

Scheme
![image](https://github.com/Rryowa/GoJira-project-manager/assets/80339180/67918f71-8604-41ec-95c6-0fffad8d41d7)

Dependencies
![image](https://github.com/Rryowa/GoJira-project-manager/assets/80339180/189e1301-f711-4429-8f71-c82e5a4bd5ea)

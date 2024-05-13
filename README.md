## **RESTful** Jira **JWT** API in **Golang**<br>
**No Web frameworks**<br>
**Layered** architecture implementation<br>
<br><br>
```***docker compose up -d***```<br>
```***make run***```<br>
**Thunder Client:**<br>
POST ```http://[::1]:3000/api/v1/users/register```<br>
with json body email, password<br>
COPY token from response(or cookies)<br>
POST ```http://[::1]:3000/api/v1/tasks```<br>
with copied token in Authorization header<br>

Scheme
![image](https://github.com/Rryowa/GoJira-project-manager/assets/80339180/67918f71-8604-41ec-95c6-0fffad8d41d7)

Dependencies
![image](https://github.com/Rryowa/GoJira-project-manager/assets/80339180/189e1301-f711-4429-8f71-c82e5a4bd5ea)

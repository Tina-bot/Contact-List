<img src= "https://64.media.tumblr.com/ce5a174b5f8b587f7f5ea4dae1f7c607/ea7d2207a691971c-f9/s500x750/e296dd4ca643bbf319e0424325ed0b98ea7f7c36.gifv">


 # Contact-List
  
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)


<img width=" 150px" align="right" src="https://raw.githubusercontent.com/golang/website/2aac4c0722c9bcb895571a7c5b68e57b45cc732a/_content/images/gophers/front.svg">

A simple list of contacts per console ready to build! <p>
Where you can **add**, **update**, **view** and **delete** the contacts that you create
  
## Annotation 
 As this program runs on console, I have a function to clean the console every time a function is executed<p>
 This function is made to run on Linux. If you run it on Windows it will run anyway, only it will print ("10") instead of cleaning the console<p>
 It's simple, you just have to change *clear* to *cls*
  
  <br>
 
 - Windows: <p>
*Line 93:*
  
   ```
   cmd := exec.Command("cls")
   ```

 - Linux: <p>
*Line 93:*
   ```
   cmd := exec.Command("clear")
   ```
   
  <br>
  
  <img width=" 150px" align="right" src="https://raw.githubusercontent.com/golang/website/2aac4c0722c9bcb895571a7c5b68e57b45cc732a/_content/images/gophers/happy.svg">

  
  ## Example output
 ```
a) Create
B) List
C) Update
D) Delete
enter an option:
c
ID: 2
New Username: Chris
New Email: ch@gmail.com
New Age: 31
User updated successfully
```

  
```
a) Create
B) List
C) Update
D) Delete
enter an option:
b
User list :
1 - Tina - tina@gmail.com - 20
2 - Chris - ch@gmail.com - 31
3 - Vivy - vio@gmail.com - 34
```

<br>

 <img src= "https://64.media.tumblr.com/ce5a174b5f8b587f7f5ea4dae1f7c607/ea7d2207a691971c-f9/s500x750/e296dd4ca643bbf319e0424325ed0b98ea7f7c36.gifv">


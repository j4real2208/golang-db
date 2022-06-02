# golang-db

A rest-API implementation with docker 

- Db mysql used 
-Logging -> zap
- Router gorilla mux
-Caching Memecached

Start Time : 23:03:40.036678 +0530 IST m=+1.714738637 
Db retrival : 23:03:40.042969 +0530 IST m=+1.721029868


Net Time Taken for db recovery of data --> .042969 - 036678 = .006291 


Start Time : 23:03:47.30516 +0530 IST m=+8.983037492 
Memcache Time : 23:03:47.305175 +0530 IST m=+8.983052910 

Net Time Taken for db recovery of data --> .305175 - .30516 = 0.000015 

On an apprrox 420x better in our current  evaluation 





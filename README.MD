#Consistent Hashing 

##Objective 

The objective of this program is to implement datastore using consistent hashing. 

The client program stores the key value in one of the three nodes using consistent hashing

The consistent hashing has the following three modules : 

1. Client
2. Server 
3. DataStores 1-3 




#How to execute ?


##Packaging :


Client.go 


Refers to the file in the folder consistent 


Refers to the file in the folder route 



Router.go 


Refers to the file in the folder router


DataStore.go/DataStore1.go/DataStore2.go 


Refers to the file in the folder server 




##Post Installation


Execute the following in separate terminals

<pre>
go run dataStore.go


go run dataStore1.go


go run dataStore2.go 
</pre>

When the dataStores are started, it will start three server instances 

<pre>
A: http://localhost:3000


B: http://localhost:3001


C: http://localhost:3002
</pre>

Each of the instances will have the capability to PUT and GET data 

<pre>
For eg: 


1.	 PUT http://localhost:3000/keys/{key_id}/{value}


2.	GET http://localhost:3000/keys/{key_id}


3.	GET http://localhost:3000/keys

go run client.go
</pre>

When the client is started it prompts the following 


Select an option


1. Execute code as a REST Application


2. Execute code without supplying the data



When option 1 is selected , 


It executes as a rest application with the following end points


1.	PUT (http://localhost:8080/keys/8/k)


Inserts the key value in one of the server instances as determined by the consistent hashing algorithm 


2.	GET(http://localhost:8080/keys/8)


Retrieves the value for the key entered from the server instances



3.	GET(http://localhost:8080/keys)


Retrieves all the values from the server instance 


When option 2 is selected, 


The predefined values are inserted in the server instances as determined by the consistent hashing algorithm


## Notes:


1.	In both of above option the values are inserted using the servers endpoint PUT 


2.	Option2 uses the end point GET from the server instances to retrieve  the data


3.	Replication factor is ignored


4.	When a server goes down, the corresponding server instance will be removed and the data will be distributed to available nodes by consistent hashing algorithm








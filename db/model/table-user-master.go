package table


// For each table, we should have a master table followed with an interface table. The master table is what lies in the database while the front end uses the interface.

type master-schema struct {
  USER-MASTER string, // This holds the true version of the user. This is the source of truth for the user. This is the table that launches the application
  USER-MASTER-INTERFACE string // This table is our interface table. This is how our client will interface with the USER_MASTER_TABLE. 
}

type user-master struct {
  USER_ID string, // PRIMARY KEY, AUTO-INCREMENT
  USERNAME string, // varChar(12), case sensitive
  EMAIL string, // varChar(24), use standard email validation
  DEVICE_NAME string, // Stored Device Name to understand users
  BROWSER_TYPE string, // Understand user browser types so I can target better
  IP_ADDRESS string, // Check to see for bad actors
}

type user-master-interface struct {
  USER_INTERFACE_ID string // PRIMARY KEY, AUTO-INCREMENT
  USER_ID string, // FOREIGN KEY, [USER_MASTER.USER_ID]
  USER_GENDER string, // For avatar box rolls
  USER_COUNTRY string, // Where is the user representing? 
  USER_BIRTHSTONE string, // What stone are you?
  USER_WERK string, // What do you do?
  USER_HOBBIES string, // Favorite hobbies. Bubble chart idea. Delete them from the grid using css animations.
  USER_BIO string, // Whatcha doing on the app?
}





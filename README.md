# Groupie-Tracker

* **Launch the project with:**
```bash 
go run .\serv\server.go
```
* **Open your browser and go to:**
```url
http://localhost:8080/
```

## Availables features
### <ins>Bonuses:</ins>
- ### Search Bar:
Options available at the top of the main webpage
Below the filters, you can search any elements of a group to find it : *Name, Members, Creation Date, First Album Date, Locations of concerts*
After clicking the button "Search", the page refreshed and displaying the corresponding groups.
Also, a "Back to Main page" button is created for reset the research.

- ### Filters:
Options available at the top the main webpage like search bar.
You can choose between 4 different filters:
- The creation date of the band/artist
- The first album release date
- The number of members
- The Locations of concerts

This filters can be used simultaneously and sort groups with multiple filters at the same time.
The filters : "*Creation Date*" and "*First Album Date*" can be disabled by the checkbox before them.
The filters must be applied by clicking on the "*filter*" button at the right of filters.
After that, the page is refreshed by displaying only the corresponding groups.
Like Research Bar, when filter applied, a "Back to Main Page" button is available to reset the filters and display all the groups. 

- ### Visualizations:
This bonus is a global improvement and optimization for all webpages.
We tried to have an easy-to-use interface, understandable by the user with work on the graphic of the pages.
Moreover, you can use the copy-past shortcut in the search bar and easily return to main page with link "*Back to Main Page*".

- ### Geolocalization:
This feature is available when you clicked on a band/artist at the bottom of the page. The map marks the differents locations of concerts.
If there are multiple markers in an area, a grouping marker is created up to user zoom to a more precise location.
Moreover, at the start, the appearance of the map is automatically located on the marks.
---
layout: post
title: "Week 7 Progress Report MiniProject"
date: 2023-03-10
---

This was quite interesting as myself along with my team of 5 members worked on a ***Library Management System*** It was a simple application that works as follows 

```
When you open the application the user login opens 

After you log in you can see books which are already there in the database

If you are an admin you can go to add book and add your own book.

On the Checkout page select the book id you need to checkout

Press the button and it checks out 

As for now you cannnot checkout the same book again

```

The front end is through Fyne package in Go whereas the backend is MYSQL

My portion was to create an ***ADD BOOK PAGE*** this page could add books to the database and then display it to the checkout page 

***The challenges faced***

The usage of package ***FYNE*** is quite complex in windows as it requires some dependencies that may or may not be present on a windows device,
once the program was compiled and when i saw the page it initially had some parameters missing although it was already specified,it took me a while to figure that out 
```

	dialog.ShowForm("Add Book", "Add", "Cancel", form, func(ok bool) {
		if ok {
			newBook := model.Book{
				Title:      titleEntry.Text,
				Author:     authorEntry.Text,
				Publisher:  publisherEntry.Text,
				ISBN:       isbnEntry.Text,
				IsBorrowed: false,
				Desc:       descEntry.Text,
			}
        }
    })

```
were the parameters and i could not get ***Desc*** running as the variable was different than that mentioned in the datamodel
but at the end i did get eveything running 






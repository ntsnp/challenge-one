import os
import requests
import json
import sqlite3
import argparse

def grab_params(page):
    index = "https://blog.sentry.io/page-data/index/page-data.json"
    if(page>1):
        apiurl = f"https://blog.sentry.io/page-data/posts/{page}/page-data.json"
    else:
        apiurl = index
    print(f'''Grabbing Page {page}''')
    result = json.loads(requests.get(apiurl).text)
    blogs = result["result"]["data"]["allContentfulBlogPost"]["edges"]
    outlist = []
    for blog in blogs:
        thumbnail = blog["node"]["metaImage"]
        if thumbnail is not None:
            thumburl = "https:"+blog["node"]["metaImage"]["file"]["url"]
        else:
            thumburl = None
        title = blog["node"]["title"]
        body = blog["node"]["body"]["childMarkdownRemark"]["html"]
        bt_list = []
        bt_list.append(title)
        bt_list.append(body)
        bt_list.append(thumburl)
        outlist.append(bt_list)

    return outlist

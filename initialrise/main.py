import os
import requests
import json
import sqlite3
import argparse
from argparse import RawTextHelpFormatter


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

def grab_all():
    print("Counting Total Pages")
    pageno = 2
    total_pages = 0
    while True:
        url = f"https://blog.sentry.io/page-data/posts/{pageno}/page-data.json"
        response = requests.get(url).text
        total_pages += 1
        pageno += 1
        if "Ah, hell" in response:
            break;

    print(f"Total pages was found to be {total_pages}")
    pagelist = []
    for p in range(1,total_pages+1):
        out = grab_params(p)
        pagelist.append(out)

    return pagelist

def download_article(foldername):
    os.mkdir(foldername)
    os.chdir(foldername)
    alldata = grab_all()
    for page in alldata:
        for blog in page:
            title = blog[0].replace(" ", "_").replace("/","").replace("\/","") + ".html"

            body = blog[1]
            thumburl = blog[2]
            bodyedit = f'''<img src="{thumburl}" alt="thumbnail">'''+body
            with open(title,"w") as article:
                article.write(bodyedit)

def savetodb(dbname):
    db = sqlite3.connect(dbname)
    dbcur = db.cursor()
    dbcur.execute('''
            CREATE TABLE blog( pageno INTEGER,
                               title TEXT,
                               html TEXT,
                               thumburl TEXT)
                               ''')
    alldata = grab_all()
    for pageno,page in enumerate(alldata):
        for blog in page:
            title = blog[0]
            body = blog[1]
            url = blog[2]
            dbcur.execute('''INSERT into blog VALUES(?,?,?,?)''',(pageno+1,title,body,url))
            db.commit()

parser = argparse.ArgumentParser(description="Get blogs from sentry.io. Use main.py -o foldername or main.py -d dbname",formatter_class=RawTextHelpFormatter)
parser.add_argument("-o","--output",type=str, help = "Output Folder")
parser.add_argument("-d","--database",type=str, help = "Output Database File")
args = parser.parse_args()
outputfolder = args.output
databasefile = args.database

def main():
    if args.output:
        download_article(outputfolder)
    elif args.database: 
        savetodb(databasefile)
    else:
        print("Type python main.py -h") 
if __name__ == '__main__':
    main()

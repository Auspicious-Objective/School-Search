import requests
from bs4 import BeautifulSoup
from lxml import etree
import pandas as pd
from concurrent.futures import ThreadPoolExecutor
from threading import Thread
import csv

#xpath data

xpathData = {'schoolNamePath' : '//*[contains(concat( " ", @class, " " ), concat( " ", "page-title", " " ))]',
             'acceptanceRatePath' : '//*[contains(concat( " ", @class, " " ), concat( " ", "field--name-field-acceptance-rate", " " ))]',
             'locationPath' : '//*[contains(concat( " ", @class, " " ), concat( " ", "college-overview--info--location", " " ))]',
             'costPath' : '//*[contains(concat( " ", @class, " " ), concat( " ", "field--name-field-average-net-price", " " ))]',
             'studentToTeacherPath' : '//*[contains(concat( " ", @class, " " ), concat( " ", "text-center", " " ))]//*[contains(concat( " ", @class, " " ), concat( " ", "data-big", " " ))]',
             'totalStudentsPath' : '//*[contains(concat( " ", @class, " " ), concat( " ", "field--name-field-total-students", " " ))]',
             'satReqPath' : '//*[contains(concat( " ", @class, " " ), concat( " ", "layout--25-25-25-25--50", " " ))]//div[(((count(preceding-sibling::*) + 1) = 1) and parent::*)]//*[contains(concat( " ", @class, " " ), concat( " ", "data-set--centered", " " ))]//*[contains(concat( " ", @class, " " ), concat( " ", "text--body-2", " " ))]'}


def makeFile():
    #makes the file and writes the headers
    headerWriter = csv.writer(open("SchoolSpreadsheet.csv", "w"))
    headerWriter.writerow(['School Name', 'Location', 'Acceptance Rate', 'Cost', 'Student to Teacher Ratio', 'Total Students', 'Is SAT Required'])

def writer(info):
    #Writes to the pre-created CSV file.
    dataFrame = pd.DataFrame(info, index=[0])

    csvConversion = dataFrame.to_csv(index=False, header=False)

    #df.to_excel(writer, startrow=rowPosition, index= False, header= False)
    with open("SchoolSpreadsheet.csv", "a") as f:
        f.write(csvConversion)
        
def getSchoolInfo(url):

    response = requests.get(url)
    soup = BeautifulSoup(response.content, "html.parser")
    global  body
    body = soup.find("body")

    parse = etree.HTML(str(body)) # Parse the HTML content of the page

    # Gets the data and  places  it in a veriable
    schoolName = parse.xpath(xpathData['schoolNamePath'])[0].text
    acceptanceRate = parse.xpath(xpathData['acceptanceRatePath'])[0].text
    location = parse.xpath(xpathData['locationPath'])[0].text
    cost = parse.xpath(xpathData['costPath'])[0].text
    studentToTeacher = parse.xpath(xpathData['studentToTeacherPath'])[0].text
    totalStudents = parse.xpath(xpathData['totalStudentsPath'])[0].text
    satReq = parse.xpath(xpathData['satReqPath'])[0].text

    global info
    info = {'name' : schoolName, 'location': location, 'acceptanceRate' : acceptanceRate,
            'cost' : cost, 'studentToTeacherRatio' : studentToTeacher, 'total' : totalStudents, 'SATreq' : satReq}

    print(f"Printing info for school: {info['name']}")
    print(info)

    writer(info)

def main(links) :
    #runs the getinfo function in threads in order to speed up the program run time
    executor = ThreadPoolExecutor(max_workers=5)

    for link in links:
        executor.submit(getSchoolInfo, link)
    
    executor.shutdown(wait=True)

def start():
    makeFileThread = Thread(target=makeFile)
    makeFileThread.start()
    global links
    links = []
    print("\nPlease use https://www.appily.com/colleges \n")
    while True:
        getInput = input("input link to school or put f for finished: ").lower()

        if getInput in links:
            print(f"'{getInput}' has already been entered")
        else:
            if getInput == "f":
                makeFileThread.join()
                main(links)
                return        
            else:
                links.append(getInput)

start()
print("\nTask completed! Enjoy your college spreadsheet :)")

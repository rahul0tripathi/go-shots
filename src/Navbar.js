import React, { useState } from 'react'
import Container from 'react-bootstrap/Container'
import Nav from 'react-bootstrap/Nav'
import Navbar from 'react-bootstrap/Navbar'
import NavDropdown from 'react-bootstrap/NavDropdown'
import Dropdown from 'react-bootstrap/Dropdown'
import Button from 'react-bootstrap/Button'
import { HiOutlinePlay } from 'react-icons/hi'
import { fetchContent } from './utils/github'
import { fmt, run } from './utils/goplay'

export function NavbarComponent({
  files,
  val,
  setVal,
  setExecState,
  execState
}) {
  const [selected, setSelected] = useState(0)
  async function changeFile(val) {
    setSelected(val)
    if (files[selected].fileName.endsWith('.go')) {
      const pathURI = files[selected].resourceURI
      setVal(await fetchContent(pathURI))
    }
  }
  async function runCode() {
    setExecState(['started execution'])
    const formatted = await fmt(val)
    if (formatted.err == '') {
      setVal(formatted.body)
      setExecState((execState) => [...execState, 'formatted code with fmt'])
      setExecState((execState) => [...execState, 'execution started'])
      const response = await run(formatted.body)

      if (response.err == '') {
        console.log(execState)
        for (let event of response.events) {
          setExecState((execState) => [
            ...execState,
            `${event?.Kind}:   ${event?.Message}`
          ])
        }
      } else {
        setExecState((execState) => [
          ...execState,
          'failed to execute',
          response.err
        ])
      }
    } else {
      setExecState((execState) => [
        ...execState,
        'failed to execute',
        formatted.err
      ])
    }
  }
  return (
    <>
      <Navbar bg="dark" variant="dark">
        <Container fluid>
          <Navbar.Brand href="#home">go-shots</Navbar.Brand>
          <Nav className="me-auto">
            <NavDropdown
              title={files[selected].fileName.slice(0, 13)}
              id="collasible-nav-dropdown"
              menuVariant="dark"
            >
              {files.map((v, i) => {
                return (
                  <>
                    <NavDropdown.Item
                      key={i}
                      onClick={() => {
                        changeFile(i)
                      }}
                    >
                      {v.fileName}
                    </NavDropdown.Item>{' '}
                    <NavDropdown.Divider />{' '}
                  </>
                )
              })}
            </NavDropdown>
          </Nav>
          <Button
            variant="danger"
            className="d-flex"
            style={{ padding: '10px 40px 10px 40px ' }}
            onClick={runCode}
          >
            <HiOutlinePlay
              style={{ fontSize: '25px', verticalAlign: 'middle' }}
            />{' '}
            <div style={{ verticalAlign: 'middle' }}>Run</div>
          </Button>
        </Container>
      </Navbar>
    </>
  )
}

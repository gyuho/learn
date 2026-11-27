# Electrical engineering roadmap

Add **Electrical engineering** as a peer of **Consensus systems** on gyuho.dev. Publish eight articles that connect electrical trade fundamentals to industrial controls and robotics. Work through the stages below in order. Use Part 1 to establish the format, and update the exam coverage map as each draft takes shape.

The supplied outline gives detailed topics for Parts 1–7 and mentions Part 8 without describing it. Part 8 is a proposal based on wiring and installation topics in the published Hawaiʻi examination outlines. The initial targets are Journey Worker Electrician (EJ) and Journey Worker Industrial Electrician (EJI).

## 1. Set up the article format and exam references

- [ ] Use the [Consensus article layout](sites/consensus-systems/paxos-etcd-vs-nakamoto-bitcoin/index.html). Keep the article contents menu, section anchors, light and dark themes, and previous and next article links.
- [ ] Put each article at `sites/electrical-engineering/<slug>/index.html`. Keep its local figures in an `img/` directory beside the article.
- [ ] Set up the Part 1 draft with prerequisites, a physical model, an original diagram, a trade example with units and assumptions, a machine or robot application where useful, a practice question with a solution, and references. Use this format for the remaining articles.
- [ ] Sketch a fictional industrial robot cell to reuse in examples on distribution, motion, control, and stored energy.
- [ ] Confirm the current PSI bulletin and exam reference edition for EJ and EJI through the official Hawaii Trade portal. Use the March 22, 2024 bulletin as the starting point for the coverage map below. [Hawaiʻi DCCA exam information](https://cca.hawaii.gov/pvl/boards/electrician/application-deadline-examination-dates/) · [PSI Hawaii Trade portal](https://test-takers.psiexams.com/hitrade) · [PSI electrician candidate bulletin, March 2024](https://test-takers.psiexams.com/api/content/bulletin/10695)

## 2. Draft the articles in order

Article slugs are relative to `/electrical-engineering/`.

### Part 1 — Three-phase AC, DC power electronics, and circuit physics

`power-and-circuit-physics`

- [ ] Explain voltage, current, impedance, inductive and capacitive reactance, RMS, and phasors. Derive the balanced three-phase √3 relationship.
- [ ] Compare line-to-line and line-to-neutral voltages for 120/208 V wye, 277/480 V wye, and 120/240 V high-leg delta. Draw and label the wye and high-leg delta phasors.
- [ ] Relate real, reactive, and apparent power (kW, kVAR, and kVA) and power factor to conductor heating and voltage drop.
- [ ] Explain 24 V and 48 V robot buses, conceptual 400 V architectures, buck and boost converters, battery C-rate and discharge curves, peak versus continuous current, and regenerative braking. Relate dynamic current to conductor heating. Simulate the bus through acceleration and regeneration.

### Part 2 — Grounding, bonding, EMI, and fault-current paths

`grounding-bonding-and-emi`

- [ ] Trace normal and fault currents on a service-to-load one-line diagram. Identify the grounded conductor, equipment grounding conductor (EGC), and grounding electrode conductor (GEC).
- [ ] Explain the main bonding jumper and bonding in separately derived systems. Check where bonding occurs for each transformer or generator topology, including the transfer arrangement.
- [ ] Compare common-mode noise on a motor and encoder harness. Show how motor returns, 3.3 V and 5 V logic references, star grounding, twisted differential pairs, and cable shields affect electromagnetic interference (EMI). Check shield termination against equipment guidance and the frequencies involved.

### Part 3 — Overcurrent protection and machine safety

`overcurrent-protection-and-machine-safety`

- [ ] Distinguish overloads, short circuits, and ground faults. Read a pair of time-current curves and explain selective coordination.
- [ ] Compare available fault current with interrupting ratings and short-circuit current ratings (SCCR). Relate those distinctions to fuse and breaker coordination in battery systems.
- [ ] Use ISO 13849 and NFPA 79 as sources for machinery examples. Explain dual-channel emergency stops, light curtains, and safe torque off (STO). Trace stop, isolation, and restart as separate states in a machine safety diagram. Verify STO behavior against the example drive's documentation.

### Part 4 — Transformers, distribution, and harmonics

`transformers-distribution-and-harmonics`

- [ ] Follow a utility-to-cell one-line diagram through transformers, switchgear, panelboards, and current transformer (CT) metering. Explain turns ratio, open delta, buck–boost transformers, and the open-CT-secondary hazard.
- [ ] Explain how delta–wye transformers establish a neutral and how their connection affects zero-sequence and triplen currents. Compare third, fifth, and seventh harmonics and neutral current for example loads.
- [ ] Connect switch-mode power supply inrush and harmonics from nonlinear loads to control power and cell distribution.

### Part 5 — Motors, drives, and motion control

`motors-drives-and-motion-control`

- [ ] Explain induction motor slip, synchronous speed, locked-rotor current, and 9-lead and 12-lead dual-voltage winding diagrams. Use the synchronous-speed formula 120f/P and calculate slip.
- [ ] Trace two- and three-wire start/stop ladder circuits using NEMA and IEC conventions. Cover holding contacts, forward/reverse interlocks, and timers.
- [ ] Describe variable-frequency drive (VFD) rectifiers, DC buses, IGBT inverters, PWM, and reflected-wave voltage spikes.
- [ ] Trace a servo axis command through torque, velocity, and position loops. Relate the example to BLDC/PMSM motors, field-oriented control, servo and stepper drives, and electromechanical holding brakes.

### Part 6 — Instrumentation, PLCs, and industrial networks

`instrumentation-plcs-and-networks`

- [ ] Trace a simulated 4–20 mA transmitter from loop current to PLC engineering units. Explain noise immunity and how live zero helps detect a broken wire. Introduce RTDs, thermocouples, and pressure and flow transmitters.
- [ ] Explain sinking (NPN) and sourcing (PNP) I/O, PLC scan cycles, and closed-loop PID tuning.
- [ ] Distinguish RS-485 from Modbus RTU and describe the role of Modbus TCP, BACnet, EtherNet/IP, EtherCAT, and CAN/CANopen.
- [ ] Map rotary encoders, Hall-effect limit switches, force-torque sensors, and IMUs to a ROS 2 and controller architecture. Show where cycle time and latency affect the robot control loop.

### Part 7 — Critical power, energy storage, and electrical safety

`critical-power-and-energy-storage`

- [ ] Compare power paths before and after a utility outage. Explain automatic transfer switches, generator paralleling and synchronization in voltage, frequency, and phase, and double-conversion UPS systems.
- [ ] Cover solar PV, battery energy storage, and AC and DC arc behavior. Explain NFPA 70E incident energy and arc-flash boundaries, lockout/tagout, and live-dead-live verification.
- [ ] Explain battery management system supervision, DC contactors, and the role of pre-charge in managing stored energy in robot drives. Plot bus voltage during a simulated pre-charge cycle.

### Part 8 — Wiring methods, code navigation, and exam practice (proposed)

`wiring-methods-and-exam-practice`

- [ ] Use a fictional service-to-load drawing to practice code lookup for services, feeders, branch circuits, conductor ampacity and adjustment, cables, raceways, and boxes.
- [ ] Add exercises for special occupancies, conditions, and equipment, low-voltage circuits, and lighting. Link relevant exercises to grounding, protection, motors, and safety in the earlier articles.
- [ ] Write mixed practice sets for EJ and EJI using their respective outlines. For each problem, show where to find the rule, which inputs determine the result, and how to check the answer.

## 3. Complete the exam coverage map and technical review

Check off a topic when each confirmed outline item links to an article section and a worked problem. Treat PLCs, industrial networks, robotics, and commissioning examples as applications of the fundamentals.

- [ ] **General electrical knowledge** — Part 1, with applications in Parts 3–5.
- [ ] **Services, feeders, and branch circuits** — Parts 3, 4, and 8.
- [ ] **Grounding and bonding** — Parts 2 and 8.
- [ ] **Conductors, cables, raceways, and boxes** — Part 8, building on the voltage-drop and heating examples in Part 1.
- [ ] **Special occupancies, conditions, and equipment** — Part 8, with related power systems in Part 7.
- [ ] **Electrical power, motors, and general-use equipment** — Parts 3–5.
- [ ] **Low-voltage circuits and communications where listed** — Part 8, with applications in Part 6.
- [ ] **Lighting, listed in the EJ outline** — Part 8.
- [ ] **Safety** — Parts 2, 3, and 7, with exam practice in Part 8.

### Technical review

- [ ] Verify the field code edition and local amendments for installation claims. Cite the applicable edition and section beside each code-dependent example.
- [ ] Check technical claims against authoritative sources. Use diagrams, calculations, and simulations for hazardous power and machine safety examples.
- [ ] Have a qualified reviewer check code and safety examples before publishing the articles that contain them.
- [ ] Complete the coverage map before describing the scope of the series as exam preparation.

## 4. Add site navigation and publish each ready article

- [ ] Add the **Electrical engineering** section after **Consensus systems** on the [homepage](sites/index.html). Use the anchor `/#electrical-engineering`, include it in the homepage contents, and list published articles in reading order.
- [ ] Add the series to the sidebar across Home, About me, Blog, the Consensus articles, and the Electrical engineering articles. Link the parent to the homepage section and nest the published articles beneath it. Use the existing current-page styling and `aria-current="page"` on each active article link.
- [ ] For each article ready to publish, add its canonical and Open Graph metadata, update the [sitemap](sites/sitemap.xml), and add its navigation entries in the same change.
- [ ] Verify desktop and mobile navigation, keyboard access, light and dark themes, figure legibility and alternative text, section anchors, and previous and next links. Check that canonical URLs match sitemap entries.

# Electrical engineering series plan

Add **Electrical engineering** as a peer of **Consensus systems** on gyuho.dev. Use eight articles to connect circuit physics and electrical trade fundamentals to industrial controls and robotics. Each article starts from a physical model, works through a trade example, and applies the same principle to a machine or robot where the connection is useful.

The supplied outline gives detailed topics for Parts 1–7 and mentions Part 8 without describing it. Part 8 below is a proposal informed by the wiring and installation topics in the published Hawaiʻi PSI exam outlines. Use Journey Worker Electrician (EJ) and Journey Worker Industrial Electrician (EJI) as the initial exam references.

## Site structure

- In [the homepage](sites/index.html), add a section at `/#electrical-engineering` after Consensus systems. List the published articles in reading order and add the section to the homepage contents.
- Follow the existing sidebar pattern across Home, About me, Blog, the Consensus articles, and the new articles. Point the Electrical engineering parent link to the homepage section and nest article links beneath it. Use the existing current-page styling and `aria-current="page"` on the active article.
- Put each article at `sites/electrical-engineering/<slug>/index.html`, with any local figures in its `img/` directory. Use the [Consensus article layout](sites/consensus-systems/paxos-etcd-vs-nakamoto-bitcoin/index.html) for the article contents menu, section anchors, light and dark themes, and adjacent article links.
- Give each published route canonical and Open Graph metadata and add it to [the sitemap](sites/sitemap.xml). Publish navigation entries as their articles become available.

## Article sequence

The slugs below are relative to `/electrical-engineering/`.

| Part | Proposed article and slug | Coverage | Worked example or figure |
| --- | --- | --- | --- |
| 1 | **Three-phase AC, DC power electronics, and circuit physics** · `power-and-circuit-physics` | Start with voltage, current, impedance, reactance, RMS, and phasors. Derive the balanced three-phase √3 relationship. Compare 120/208 V wye, 277/480 V wye, and 120/240 V high-leg delta. Relate real, reactive, and apparent power to voltage drop and conductor heating. Extend to 24 V and 48 V robot buses and conceptual 400 V architectures, buck and boost converters, battery C-rate, peak versus continuous current, and regenerative braking. | Annotated wye and high-leg delta phasors, then a simulated robot bus through acceleration and regeneration. |
| 2 | **Grounding, bonding, EMI, and fault-current paths** · `grounding-bonding-and-emi` | Trace the grounded conductor, equipment grounding conductor, and grounding electrode conductor in normal and fault conditions. Explain bonding at services and separately derived systems, including generator topology. Relate motor return currents and logic references to differential pairs and shielding. | Trace the current paths on a service-to-load one-line and compare common-mode noise on a motor and encoder harness. |
| 3 | **Overcurrent protection and machine safety** · `overcurrent-protection-and-machine-safety` | Cover overloads, short circuits, ground faults, time-current curves, coordination, available fault current, interrupting rating, and short-circuit current rating (SCCR). Introduce battery protection and the roles of dual-channel E-stops, light curtains, and safe torque off (STO). Use ISO 13849 and NFPA 79 as sources for the machinery examples. | Read a pair of time-current curves and trace separate stop, isolation, and restart states in a machine safety diagram. |
| 4 | **Transformers, distribution, and harmonics** · `transformers-distribution-and-harmonics` | Cover turns ratio, delta–wye and zero-sequence behavior, open delta, buck–boost, switchgear, panelboards, and current transformer (CT) metering. Explain the open-secondary hazard conceptually. Connect switch-mode supply inrush and nonlinear loads to cell power distribution and harmonic currents. | Follow a utility-to-cell one-line, then compare current spectra and neutral current for example loads. |
| 5 | **Motors, drives, and motion control** · `motors-drives-and-motion-control` | Cover induction motors, slip, synchronous speed, locked-rotor current, and dual-voltage winding diagrams. Introduce two- and three-wire ladder circuits, holding contacts, interlocks, timers, VFDs, PWM, and reflected waves. Extend to BLDC/PMSM, field-oriented control, servo and stepper drives, nested control loops, and holding brakes. | Compute synchronous speed and slip, then follow a start/stop ladder sequence and a servo axis command through its control loops. |
| 6 | **Instrumentation, PLCs, and industrial networks** · `instrumentation-plcs-and-networks` | Cover 4–20 mA scaling and live zero, sourcing and sinking I/O, RTDs, thermocouples, pressure and flow transmitters, PLC scans, and PID. Distinguish RS-485 from Modbus RTU, then place Modbus TCP, BACnet, EtherNet/IP, EtherCAT, and CAN/CANopen in context. Map encoders, Hall switches, force sensors, and IMUs to a ROS 2 and controller architecture. | Trace one simulated transmitter from current to PLC engineering units, then show where cycle time and latency enter a robot control loop. |
| 7 | **Critical power, energy storage, and electrical safety** · `critical-power-and-energy-storage` | Cover transfer switches, generator synchronization, double-conversion UPS, PV and battery storage, AC and DC arc behavior, arc-flash boundary concepts, and lockout/tagout. Relate BMS supervision, DC contactors, and pre-charge to stored energy in robot drives. | Compare the power paths before and after a utility outage and plot bus voltage during a simulated pre-charge cycle. |
| 8 | **Wiring methods, code navigation, and exam practice** · `wiring-methods-and-exam-practice` · proposed | Practice finding requirements for services, feeders, branch circuits, conductor ampacity and adjustment, cables, raceways, boxes, special occupancies and equipment, low-voltage circuits, and lighting. Pull together applicable grounding, protection, motor, and safety topics from the earlier parts. | Work from a fictional service-to-load drawing and show where to find each applicable rule, which inputs determine the result, and how to check the answer. Include mixed practice for EJ and EJI according to their respective outlines. |

## Exam coverage

Use the publicly available PSI bulletin dated March 22, 2024 as an initial outline. Before publication, verify the bulletin and exam reference edition for each target classification through the official Hawaii Trade portal. Verify the field code and applicable local amendments separately when an example makes an installation claim. Cite the edition and section beside each code-dependent example.

| Topic in the published outlines | Planned coverage |
| --- | --- |
| General electrical knowledge | Part 1, with applied examples in Parts 3–5 |
| Services, feeders, and branch circuits | Parts 3, 4, and 8 |
| Grounding and bonding | Parts 2 and 8 |
| Conductors, cables, raceways, and boxes | Part 8, building on the voltage-drop and heating examples in Part 1 |
| Special occupancies, conditions, and equipment | Part 8, with related power systems in Part 7 |
| Electrical power, motors, and general-use equipment | Parts 3–5 |
| Low-voltage circuits and communications where listed | Part 8, with industrial applications in Part 6 |
| Lighting, listed in the EJ outline | Part 8 |
| Safety | Parts 2, 3, and 7, with exam practice in Part 8 |

Map each confirmed outline item to an article section and a worked problem before describing exam coverage. Treat PLCs, industrial networks, robotics, and commissioning examples as applications of the fundamentals. [Hawaiʻi DCCA exam information](https://cca.hawaii.gov/pvl/boards/electrician/application-deadline-examination-dates/) · [PSI Hawaii Trade portal](https://test-takers.psiexams.com/hitrade) · [PSI electrician candidate bulletin, March 2024](https://test-takers.psiexams.com/api/content/bulletin/10695)

## Drafting and review

1. Draft Part 1 first and use it to establish the article format. Include prerequisites, an original diagram, a worked example with units and assumptions, a practice question with a solution, and references. Reuse a fictional industrial robot cell across the series where it helps readers connect distribution, motion, control, and stored energy.
2. Draft Parts 2–5 in order, then Parts 6–7. Build the Part 8 exam map alongside the drafts so code-lookup exercises can link back to the relevant explanations.
3. Review technical claims against authoritative sources. Check bonding location by system and transfer topology, shield termination by equipment and frequency, available fault current against equipment ratings, and STO behavior against the drive documentation. Cover hazardous power and machine safety through diagrams, calculations, and simulation. Have a qualified reviewer check code and safety examples before publication.
4. Verify the homepage and sidebar hierarchy in desktop and mobile layouts. Check keyboard navigation, light and dark themes, figure text and alternative text, section anchors, previous and next links, and consistency among canonical URLs and sitemap entries.
